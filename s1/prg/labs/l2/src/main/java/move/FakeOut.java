package move;

import ru.ifmo.se.pokemon.Effect;
import ru.ifmo.se.pokemon.PhysicalMove;
import ru.ifmo.se.pokemon.Pokemon;
import ru.ifmo.se.pokemon.Type;

final public class FakeOut extends PhysicalMove {
  public FakeOut() {
    super(Type.NORMAL, 40, 1, 3, 1);
  }

  @Override
  public void applyOppEffects(Pokemon p) {
    Effect.flinch(p);
  }

  @Override
  public String describe() {
    return "uses ability \"Fake Out\"";
  }
}
