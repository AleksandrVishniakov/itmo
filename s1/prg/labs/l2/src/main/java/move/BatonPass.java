package move;

import ru.ifmo.se.pokemon.StatusMove;
import ru.ifmo.se.pokemon.Type;

final public class BatonPass extends StatusMove {
  public BatonPass() {
    super(Type.NORMAL, 0, 1);
  }

  @Override
  public String describe() {
    return "uses ability \"Baton Pass\"";
  }
}
