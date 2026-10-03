package move;

import ru.ifmo.se.pokemon.Effect;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.SpecialMove;
import ru.ifmo.se.pokemon.Stat;
import ru.ifmo.se.pokemon.Type;

final public class MuddyWater extends SpecialMove {
  public MuddyWater() {
    super(Type.WATER, 90, 0.85);
  }

  @Override
  public void applyOppEffects(Pokemon p1) {
    if (Math.random() <= 0.3) {
      System.out.println("Opponent's accuracy decreased!");
      p1.addEffect(new Effect().stat(Stat.ACCURACY, -1));
    }
  }

  @Override
  public String describe() {
    return "uses ability \"Muddy Water\"";
  }
}
