package move;

import ru.ifmo.se.pokemon.Effect;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Stat;
import ru.ifmo.se.pokemon.StatusMove;
import ru.ifmo.se.pokemon.Type;

final public class Flatter extends StatusMove {
  public Flatter() {
    super(Type.DARK, 0, 1);
  }

  @Override
  public void applyOppEffects(Pokemon p) {
    Effect.confuse(p);
    p.addEffect(new Effect().stat(Stat.SPECIAL_ATTACK, 1));
  }

  @Override
  public String describe() {
    return "uses ability \"Flatter\"";
  }
}
